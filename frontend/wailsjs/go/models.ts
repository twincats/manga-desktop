export namespace download {
	
	export class FiledChapReport {
	    chap_id: string;
	    chapter: string;
	    status_db: boolean;
	    error: any;
	    fail_page: internet.StatDownload[];
	
	    static createFrom(source: any = {}) {
	        return new FiledChapReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chap_id = source["chap_id"];
	        this.chapter = source["chapter"];
	        this.status_db = source["status_db"];
	        this.error = source["error"];
	        this.fail_page = this.convertValues(source["fail_page"], internet.StatDownload);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PageReport {
	    manga: string;
	    status_dl: boolean;
	    error: string;
	    fail_chap: FiledChapReport[];
	
	    static createFrom(source: any = {}) {
	        return new PageReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = source["manga"];
	        this.status_dl = source["status_dl"];
	        this.error = source["error"];
	        this.fail_chap = this.convertValues(source["fail_chap"], FiledChapReport);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ParamJS {
	    manga_id: number;
	    manga: string;
	    chapter: types.ChapterList;
	    pages: string[];
	
	    static createFrom(source: any = {}) {
	        return new ParamJS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga_id = source["manga_id"];
	        this.manga = source["manga"];
	        this.chapter = source["chapter"];
	        this.pages = source["pages"];
	    }
	}

}

export namespace file {
	
	export class Convert {
	    title: string;
	    quality: number;
	    parallel: number;
	    resize_width: number;
	    resize: boolean;
	    delete: boolean;
	    engine: boolean;
	    ext: string[];
	
	    static createFrom(source: any = {}) {
	        return new Convert(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.quality = source["quality"];
	        this.parallel = source["parallel"];
	        this.resize_width = source["resize_width"];
	        this.resize = source["resize"];
	        this.delete = source["delete"];
	        this.engine = source["engine"];
	        this.ext = source["ext"];
	    }
	}
	export class ConvertResult {
	    manga: string;
	    size_before: number;
	    size_after: number;
	    size_percent: number;
	    total_ok: number;
	    total_failed: number;
	    total_resize: number;
	    total_delete: number;
	
	    static createFrom(source: any = {}) {
	        return new ConvertResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = source["manga"];
	        this.size_before = source["size_before"];
	        this.size_after = source["size_after"];
	        this.size_percent = source["size_percent"];
	        this.total_ok = source["total_ok"];
	        this.total_failed = source["total_failed"];
	        this.total_resize = source["total_resize"];
	        this.total_delete = source["total_delete"];
	    }
	}

}

export namespace internet {
	
	export class StatDownload {
	    index: number;
	    status_resume: boolean;
	    filename: string;
	    name: string;
	    url: string;
	    err: any;
	
	    static createFrom(source: any = {}) {
	        return new StatDownload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.status_resume = source["status_resume"];
	        this.filename = source["filename"];
	        this.name = source["name"];
	        this.url = source["url"];
	        this.err = source["err"];
	    }
	}

}

export namespace manga {
	
	export class Alter {
	    id: number;
	    title: string;
	    manga_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Alter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.manga_id = source["manga_id"];
	    }
	}
	export class Chapter {
	    id: number;
	    chapter: number;
	    title: string;
	    volume: string;
	    group: string;
	    timestamp_release: number;
	    status_read: boolean;
	    language_id: number;
	    language: Language;
	    created_at: Date;
	    updated_at: Date;
	
	    static createFrom(source: any = {}) {
	        return new Chapter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.chapter = source["chapter"];
	        this.title = source["title"];
	        this.volume = source["volume"];
	        this.group = source["group"];
	        this.timestamp_release = source["timestamp_release"];
	        this.status_read = source["status_read"];
	        this.language_id = source["language_id"];
	        this.language = source["language"];
	        this.created_at = new Date(source["created_at"]);
	        this.updated_at = new Date(source["updated_at"]);
	    }
	}
	export class Config {
	    id: number;
	    manga_folder: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.manga_folder = source["manga_folder"];
	    }
	}
	export class Language {
	    id: number;
	    lang: string;
	    lang_code: string;
	
	    static createFrom(source: any = {}) {
	        return new Language(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.lang = source["lang"];
	        this.lang_code = source["lang_code"];
	    }
	}
	export class Manga {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex: string;
	    chapter: Chapter[];
	    alter: Alter[];
	
	    static createFrom(source: any = {}) {
	        return new Manga(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
	        this.chapter = this.convertValues(source["chapter"], Chapter);
	        this.alter = this.convertValues(source["alter"], Alter);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MangaHomeApi {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex: string;
	    chapter_id: number;
	    chapter: number;
	    download_time: Date;
	
	    static createFrom(source: any = {}) {
	        return new MangaHomeApi(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
	        this.chapter_id = source["chapter_id"];
	        this.chapter = source["chapter"];
	        this.download_time = new Date(source["download_time"]);
	    }
	}
	export class MangaHome {
	    manga: MangaHomeApi[];
	    pagination?: Pagination;
	
	    static createFrom(source: any = {}) {
	        return new MangaHome(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = this.convertValues(source["manga"], MangaHomeApi);
	        this.pagination = source["pagination"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Page {
	    pages: string[];
	    path: string;
	    chaps: number[];
	    manga_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Page(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pages = source["pages"];
	        this.path = source["path"];
	        this.chaps = source["chaps"];
	        this.manga_id = source["manga_id"];
	    }
	}
	export class Pagination {
	    current_page: number;
	    from: number;
	    to: number;
	    total: number;
	    last_page: number;
	    perpage: number;
	
	    static createFrom(source: any = {}) {
	        return new Pagination(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current_page = source["current_page"];
	        this.from = source["from"];
	        this.to = source["to"];
	        this.total = source["total"];
	        this.last_page = source["last_page"];
	        this.perpage = source["perpage"];
	    }
	}
	export class Server {
	    id: number;
	    name: string;
	    url: string;
	    search: boolean;
	    status_active: boolean;
	    chap_jscode: string;
	    page_jscode: string;
	
	    static createFrom(source: any = {}) {
	        return new Server(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.url = source["url"];
	        this.search = source["search"];
	        this.status_active = source["status_active"];
	        this.chap_jscode = source["chap_jscode"];
	        this.page_jscode = source["page_jscode"];
	    }
	}

}

export namespace tool {
	
	export class Web {
	    url: string;
	    body: string;
	    error: any;
	
	    static createFrom(source: any = {}) {
	        return new Web(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.body = source["body"];
	        this.error = source["error"];
	    }
	}

}

export namespace types {
	
	export class ChapterList {
	    id: string;
	    chapter: string;
	    title: string;
	    volume: string;
	    timestamp: number;
	    language: string;
	    group_name: string;
	    status: boolean;
	    check: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ChapterList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.chapter = source["chapter"];
	        this.title = source["title"];
	        this.volume = source["volume"];
	        this.timestamp = source["timestamp"];
	        this.language = source["language"];
	        this.group_name = source["group_name"];
	        this.status = source["status"];
	        this.check = source["check"];
	    }
	}
	export class Chapter {
	    manga: string;
	    cover_url: string;
	    mdex?: string;
	    manga_id: number;
	    server_name: string;
	    chapter: ChapterList[];
	    datasaver: boolean;
	    limit: number;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new Chapter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = source["manga"];
	        this.cover_url = source["cover_url"];
	        this.mdex = source["mdex"];
	        this.manga_id = source["manga_id"];
	        this.server_name = source["server_name"];
	        this.chapter = this.convertValues(source["chapter"], ChapterList);
	        this.datasaver = source["datasaver"];
	        this.limit = source["limit"];
	        this.total = source["total"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Option {
	    url: string;
	    server_name: string;
	    offset?: number;
	    limit?: number;
	    datasaver?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Option(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.server_name = source["server_name"];
	        this.offset = source["offset"];
	        this.limit = source["limit"];
	        this.datasaver = source["datasaver"];
	    }
	}

}

