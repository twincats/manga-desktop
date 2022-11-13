export namespace manga {
	
	export class Manga {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex?: number;
	
	    static createFrom(source: any = {}) {
	        return new Manga(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
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
	
	export class Nav {
	    next: string;
	    prev: string;
	
	    static createFrom(source: any = {}) {
	        return new Nav(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.next = source["next"];
	        this.prev = source["prev"];
	    }
	}
	export class Page {
	    pages: string[];
	    path: string;
	    nav: Nav;
	
	    static createFrom(source: any = {}) {
	        return new Page(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pages = source["pages"];
	        this.path = source["path"];
	        this.nav = this.convertValues(source["nav"], Nav);
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

}

