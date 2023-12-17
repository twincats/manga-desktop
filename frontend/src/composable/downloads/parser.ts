import { types } from '@wails/go/models'
import { Fetch, FetchPost } from '@wails/go/tool/Web'

/** @type {string | null} for param Element*/
export type SNull = string | null

// class Parser
export class Parser {
  #p: DOMParser
  dom: Document
  node: NodeListOf<Element> | null = null

  /** constructor class Parser
   * @param {string} s html string
   */
  constructor(s: string) {
    this.#p = new DOMParser()
    this.dom = this.#p.parseFromString(s, 'text/html')
  }

  /** #setV is to get string from
   * @param {SNull} v html string
   */
  #setV(v: SNull): string {
    if (v != null) {
      return v.trim()
    } else {
      return ''
    }
  }

  /**
   * @param {string} s string class or element to search
   * @returns {Parser} instance of current class Parser
   */
  domFind(s: string): Parser {
    this.node = this.dom.querySelectorAll(s)
    return this
  }

  /**
   * @param {Element} el Element to search
   * @param {string} s string class or element to search
   * @returns {Element} element founded in search
   */
  elFind(el: Element, s: string): Element {
    const nel = el.querySelector(s)
    if (nel) {
      return nel
    } else {
      return el
    }
  }

  /**
   * Global forEach to loop
   * @param {Element} el Element to loop
   * @param {function } fn function callback
   * @returns {any[]} element founded in search
   */
  forEach(fn: { (el: Element, i: number): object | string | null }): any[] {
    if (this.node) {
      const o: any[] = []
      for (let x of this.node.entries()) {
        const z = fn(x[1], x[0])
        if (z != null) {
          o.push(z)
        }
      }
      return o
    } else {
      return []
    }
  }

  /**
   * forEachChap is forEach to loop return types.ChapterList
   * @param {Element} el Element to loop
   * @param {function } fn function callback
   * @returns {any[]} types.ChapterList[]
   */
  forEachChap(fn: {
    (el: Element, i: number): types.ChapterList | null
  }): types.ChapterList[] {
    if (this.node) {
      const o: types.ChapterList[] = []
      for (let x of this.node.entries()) {
        const z = fn(x[1], x[0])
        if (z != null) {
          o.push(z)
        }
      }
      return o
    } else {
      return []
    }
  }

  /**
   * getText get textContent
   * @returns {string | null} textContent of element
   */
  getText(): string | null {
    if (this.node != null && this.node.length > 0) {
      return this.node[0].textContent
    } else {
      return null
    }
  }

  /** getAttr get getAttribute()
   * @param {string} att Attribute to get
   * @returns {string | null} value of getAttribute(att)
   */
  getAttr(att: string): string | null {
    if (this.node != null && this.node.length > 0) {
      return this.node[0].getAttribute(att)
    } else {
      return null
    }
  }

  /** peek
   * @param {Element} el HTML Element
   * @param {string} s search peek string
   * @returns {boolean} return true if found
   */
  peek(el: Element, s: string): boolean {
    return el.innerHTML.startsWith(s, 0)
  }

  /**
   * @param {string | null} v unformatted chapter string
   * @returns {number} formatted chapter as number
   */
  chap(v: string | null): number {
    const x = this.#setV(v)
    const regex =
      /^(?:[Cc]h(?:apter|\.?) ?\s+?)?(?:0?)+((\d+)[\.\-,](\d+)|\d+)/gm
    var o = 0
    for (const m of x.matchAll(regex)) {
      if (m[2] && m[3]) {
        o = parseFloat(m[2] + '.' + m[3])
      } else if (m[1]) {
        o = parseFloat(m[1])
      } else {
        o = 0
      }
    }
    return o
  }
}

// class MangaParser
export class MangaParser {
  /**
   * func main dynamic func
   * @type {Function} func holding generated text func
   */
  func: Function | null = null
  funcSync: Function | null = null

  /**
   * chapter is types.Chapter
   * @type {types.Chapter} chapter for holding get chapter
   */
  chapter: types.Chapter

  /**
   * urlFetch
   * @type {string | null } holind temporary urlFetch
   */
  urlFetch: string | null = null

  /**
   * urlFetchPost
   * @type {string | null } holind temporary urlFetchPost
   */
  urlFetchPost: string | null = null

  /**
   * pages is list of url image
   * @type {types.Chapter} pages for holding get pages
   */
  pages: string[] = []

  #tmpHTMLFetch: string = ''
  #tmpHTMLFetchPost: string = ''

  constructor() {
    this.chapter = this.#setDefault()
  }

  /**
   * #setFuncA() is private for Getting Func
   * @return {Function} dynamic Function
   */
  #setFuncA(c: string, ...args: any): Function {
    return new Function(args, c)
  }

  #setFuncB(c: string, ...args: any): Function {
    const AsyncFunction = <FunctionConstructor>async function () {}.constructor
    return new AsyncFunction(args, c)
  }

  /**
   * #setDefault() is private for setting new Default Chapter
   * @return {types.Chapter} Chapter
   */
  #setDefault(): types.Chapter {
    return new types.Chapter()
  }

  /**
   * setV() is set string null to string with trim()
   * @return {string} string
   */
  setV(v: string | null): string {
    if (v != null) {
      return v.trim()
    } else {
      return ''
    }
  }

  /**
   * setT() is set string null to number unix timestamp
   * @return {number} unix timestamp : number
   */
  setT(v: string | null): number {
    if (v != null) {
      const d = Date.parse(v)
      if (isNaN(d)) {
        return Math.floor(Date.now() / 1000)
      } else {
        return Math.floor(d / 1000)
      }
    } else {
      return 0
    }
  }

  /**
   * setManga() is set string null to chapter.manga
   * @param {string | null} v auto save Manga
   * @return {void}
   */
  setManga(v: string | null) {
    this.chapter.manga = this.setV(v)
  }

  /**
   * setCover() is set string null to chapter.cover_url
   * @param {string | null} v auto save cover_url
   * @return {void}
   */
  setCover(v: string | null) {
    this.chapter.cover_url = this.setV(v)
  }

  /**
   * setChapter() is set ChapterList[] to chapter.chapter
   * @param {ChapterList[]} v ChapterList
   * @return {void}
   */
  setChapter(v: types.ChapterList[]) {
    this.chapter.chapter = v
    this.chapter.total = v.length
  }

  /**
   * setChap() is auto ChapterList[]
   * @return {MangaParser} instance of MangaParser
   */
  setChap(
    id: SNull,
    chap: SNull,
    group: SNull,
    lang: SNull,
    tm: SNull,
    tit: SNull = '',
    vol: SNull = ''
  ): types.ChapterList {
    const c = new types.ChapterList()
    c.id = this.setV(id)
    c.chapter = this.setV(chap)
    c.title = this.setV(tit)
    c.volume = this.setV(vol)
    c.timestamp = this.setT(tm)
    c.language = this.setV(lang)
    c.group_name = this.setV(group)
    return c
  }

  setURl(s: string): MangaParser {
    if (this.urlFetch == null) {
      this.urlFetch = s
    }
    return this
  }

  setURLPost(s: string): MangaParser {
    if (this.urlFetchPost == null) {
      this.urlFetchPost = s
    }
    return this
  }

  clearURLAll(): void {
    this.urlFetchPost = null
    this.urlFetchPost = null
  }

  /**
   * getFunc() is for getting and setting main func
   * @return {Function} dynamic Function
   */
  getFunc(c: string = '', ...args: any): Function {
    if (typeof this.func == 'function' && c == '') {
      return this.func
    } else {
      return (this.func = this.#setFuncA(c, 'C', ...args))
    }
  }

  getFuncSync(c: string = '', ...args: any): Function {
    if (typeof this.funcSync == 'function' && c == '') {
      return this.funcSync
    } else {
      return (this.func = this.#setFuncB(c, 'C', ...args))
    }
  }

  /**
   * err() is for throwing new Error
   * @return {void}
   */
  err(s: string) {
    throw new Error(s)
  }

  /**
   * fetch() is for fetching HTML Req return string
   * @return {void}
   */
  fetch(u: string): Promise<string> {
    return new Promise((res, rej) => {
      if (this.urlFetch != u) {
        this.urlFetch = u
        Fetch(u)
          .then(out => {
            this.#tmpHTMLFetch = out
            res(out)
          })
          .catch(e => rej(e))
      } else {
        if (this.#tmpHTMLFetch != '') {
          res(this.#tmpHTMLFetch)
        } else {
          rej('please clear urlFetch' + this.urlFetch)
        }
      }
    })
  }

  fetchPost(u: string, postData: any = null): Promise<string> {
    return new Promise((res, rej) => {
      if (this.urlFetchPost != u) {
        this.urlFetchPost = u
        FetchPost(u, postData)
          .then(out => {
            this.#tmpHTMLFetchPost = out
            console.log('out', out.substring(0, 100))
            res(out)
          })
          .catch(e => rej(e))
      } else {
        if (this.#tmpHTMLFetchPost != '') {
          console.log('out', 'darisini')
          res(this.#tmpHTMLFetchPost)
        } else {
          rej('please clear urlFetchPost')
        }
      }
    })
  }
  getParser(s: string): Parser {
    return new Parser(s)
  }
  #getManga(): types.Chapter {
    return this.chapter
  }
  getManga(u: string, fn: string): Promise<types.Chapter> {
    return new Promise((res, rej) => {
      this.fetch(u)
        .then(out => {
          try {
            const func = this.getFuncSync(fn, 'p')
            func(this, this.getParser(out))
              .then((ro: unknown) => {
                if (ro) {
                  // console.log(ro)
                }
              })
              .catch((e: unknown) => rej(e))
              .finally(() => {
                res(this.#getManga())
              })
          } catch (e) {
            rej(e)
          }
        })
        .catch(e => {
          this.urlFetch = null
          rej(e)
        })
    })
  }
  #getPages(): string[] {
    return this.pages
  }
  getPages(u: string, fn: string): Promise<string[]> {
    return new Promise((res, rej) => {
      this.fetch(u)
        .then(out => {
          try {
            const func = this.getFuncSync(fn, 'p')
            func(this, this.getParser(out))
              .then((ro: unknown) => {
                if (ro) {
                  // console.log(ro)
                }
              })
              .catch((e: unknown) => rej(e))
              .finally(() => {
                res(this.#getPages())
              })
          } catch (e) {
            rej(e)
          }
        })
        .catch(e => {
          this.urlFetch = null
          rej(e)
        })
    })
  }

  isEmptyChapter(): boolean {
    return JSON.stringify(this.chapter) === JSON.stringify(this.#setDefault())
  }

  isEmptyPage(): boolean {
    return this.pages.length == 0
  }
}
