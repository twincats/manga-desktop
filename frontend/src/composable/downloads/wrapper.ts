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
      return GetChapter(o)
    } else {
      return new Promise((res, rej) => {
        if (s.chap_jscode == '') {
          rej(`ChapJS Code is Empty. Implement ChapJS server ${s.name}`)
          return
        }
        this.#p
          .getManga(o.url, s.chap_jscode)
          .then(r => {
            r.server_name = s.name
            CheckChapterDB(r)
              .then(rr => {
                res(rr)
              })
              .catch(() => {
                res(r)
              })
          })
          .catch(e => {
            rej(e)
          })
      })
    }
  }

  GetPage(c: types.Chapter, s: manga.Server): Promise<download.PageReport> {
    if (s.id == 1) {
      return GetPage(c)
    } else {
      return new Promise(async (res, rej) => {
        if (s.page_jscode == '') {
          rej(`PagesJS Code is Empty. Implement PagesJS server ${s.name}`)
          return
        }
        const p = new download.PageReport()
        p.fail_chap = []
        var manga_id = c.manga_id
        // save mangaDB
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
        p.manga = c.manga

        // loop download chapter
        for (let i = 0; i < c.chapter.length; i++) {
          const cc = c.chapter[i]
          //report failed chap
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
                    fc.error = 'Error some pages failed download'
                    fc.fail_page = rd
                    p.status_dl = false
                    p.fail_chap.push(fc)
                    console.log('partial downloadJS finish', fc)
                    // add error retry list files
                  }
                })
                .catch(e => {
                  // error download all pages / single chapter
                  fc.error = e
                  p.status_dl = false
                  p.fail_chap.push(fc)
                  console.log('downloadJS err', e)
                })

              // console.log(r)
            })
            .catch(e => {
              // error getting list pages
              console.log('err get pages', e)
              fc.error = e
              p.status_dl = false
              p.fail_chap.push(fc)
              this.#p.clearURLAll()
            })
        } // endloop

        // end return pageresult
        p.status_dl = true
        res(p)
        this.#p.clearURLAll()
      })
    }
  }
}
