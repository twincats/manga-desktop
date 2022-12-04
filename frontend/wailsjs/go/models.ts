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
	export class Chapter {
	    id: number;
	    chapter: string;
	    title: string;
	    volume: string;
	    group: string;
	    timestamp_release: number;
	    status_read: boolean;
	    language_id: number;
	    language: Language;
	    // Go type: sqltime.Time
	    created_at: any;
	    // Go type: sqltime.Time
	    updated_at: any;
	
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
	        this.language = this.convertValues(source["language"], Language);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
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
	
	export class Manga {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex?: number;
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
	export class MangaHomeApi {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex?: number;
	    chapter_id: number;
	    chapter: string;
	    // Go type: sqltime.Time
	    download_time: any;
	
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
	        this.download_time = this.convertValues(source["download_time"], null);
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
	export class MangaHome {
	    manga: MangaHomeApi[];
	    pagination?: Pagination;
	
	    static createFrom(source: any = {}) {
	        return new MangaHome(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = this.convertValues(source["manga"], MangaHomeApi);
	        this.pagination = this.convertValues(source["pagination"], Pagination);
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
	
	export class Server {
	    id: number;
	    name: string;
	    url: string;
	    search: boolean;
	    status_active: boolean;
	
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
	    }
	}

}

