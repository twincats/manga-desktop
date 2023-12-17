import { MangaParser } from './parser'
import { manga, types, download } from '@wails/go/models'
import {
  GetChapter,
  GetPage,
  DownloadJS,
  SaveMangaCover,
  CheckChapterDB,
} from '@wails/go/download/Download'
import type { EventChap } from '@/type/download'
import { EventsEmit } from '@wails/runtime/runtime'

export class Download {
  #p: MangaParser
  constructor() {
    this.#p = new MangaParser()
  }
  GetChap(o: types.Option, s: manga.Server): Promise<types.Chapter> {
    if (s.id == 1) {
      // GetChapter 1 is for mangadex only
      return GetChapter(o)
    } else {
      // return Promise is from Chap_JSCODE rest of server
      return new Promise((res, rej) => {
        // check if Chap_JSCODE is not empty
        if (s.chap_jscode == '') {
          rej(`ChapJS Code is Empty. Implement ChapJS server ${s.name}`)
          return
        }

        // call getManga parser
        this.#p
          .getManga(o.url, s.chap_jscode)
          .then(r => {
            // r is return value from parser : types.Chapter
            r.server_name = s.name

            // checkChapterDB is check if Manga & Chapter already in DB
            CheckChapterDB(r)
              .then(rr => {
                // rr is result after checkChapterDB : types.Chapter
                res(rr)
              })
              .catch(() => {
                // if check failed return r from parser
                res(r)
              })
          })
          .catch(e => {
            // failed get : types.Chapter from parser
            rej(e)
          })
      })
    }
  } // end getChap

  GetPage(c: types.Chapter, s: manga.Server): Promise<download.PageReport> {
    if (s.id == 1) {
      // GetPage 1 is for mangadex only
      return GetPage(c)
    } else {
      // return Promise is from Pages_JSCODE rest of server
      return new Promise(async (res, rej) => {
        // check if Pages_JSCODE is not empty
        if (s.page_jscode == '') {
          rej(`PagesJS Code is Empty. Implement PagesJS server ${s.name}`)
          return
        }

        // create new page report
        const p = new download.PageReport()
        p.fail_chap = []
        p.manga = c.manga

        var manga_id = c.manga_id

        // save mangaDB if not already in DB
        if (manga_id == 0) {
          // save manga and cover
          const ret = await SaveMangaCover(c)
            .then(id => {
              if (id != 0) {
                manga_id = id
              }
              return false
            })
            .catch(e => {
              rej(e)
              return true
            })
          // return if being rejected
          if (ret) {
            rej('error when SaveMangaCover')
            return
          }
        }

        // loop download chapter
        for (let i = 0; i < c.chapter.length; i++) {
          const cc = c.chapter[i]

          // create report failed chap every loop
          const fc = new download.FiledChapReport()
          fc.chap_id = cc.id
          fc.chapter = cc.chapter

          await this.#p
            .getPages(cc.id, s.page_jscode)
            .then(async r => {
              // do event
              const evc: EventChap = {
                chapter: parseFloat(cc.chapter),
                index_chap: i + 1,
                total_chap: c.chapter.length,
                total_page: r.length,
                list_page: r,
              }

              // emit eventchap
              EventsEmit('dl_eventchap', evc)

              // do download
              const paramJs = new download.ParamJS()
              paramJs.manga_id = c.manga_id
              paramJs.manga = c.manga
              paramJs.chapter = cc
              paramJs.pages = r
              await DownloadJS(paramJs)
                .then(rd => {
                  // finish complete
                  if (rd) {
                    if (rd.length > 0) {
                      fc.error = 'Error some pages failed download'
                      fc.fail_page = rd
                      console.log(
                        '%cErr Partial downloadJS finish : %o',
                        'color:red',
                        rd
                      )
                    }
                    // add error retry list files
                  }
                })
                .catch(e => {
                  // error download all pages / single chapter
                  fc.error = e
                  console.log('%cErr Catch downloadJS : %o', 'color:red', e)
                })

              // end of download list pages
            })
            .catch(e => {
              // error getting list pages
              console.log('%cErr Get Pages : %o', 'color:red', e)
              fc.error = e
            })

          // push filechapReport if show error
          if (fc.error) {
            p.fail_chap.push(fc)
          }
        } // endloop

        // end download return pageresult
        if (p.fail_chap.length > 0) {
          p.status_dl = false
          p.error = 'Finish Download with Error'
        } else {
          p.status_dl = true
          p.error = ''
        }

        p.manga_id = manga_id
        // return pageResult to promise
        res(p)
        this.#p.clearURLAll()
      })
    }
  } // end getpage
}
