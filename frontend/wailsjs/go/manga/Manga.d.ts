// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {manga} from '../models';

export function GetManga(arg1:number):Promise<manga.Manga>;

export function GetMangaHome(arg1:any,arg2:any,arg3:number,arg4:number):Promise<manga.MangaHome>;

export function GetMangaWithAlter(arg1:number):Promise<manga.Manga>;

export function GetMangaWithChapter(arg1:number):Promise<manga.Manga>;

export function GetMangas():Promise<Array<manga.Manga>>;

export function GetPage(arg1:number):Promise<manga.Page>;

export function GetRandomMangaHome(arg1:number):Promise<Array<manga.MangaHomeApi>>;

export function GetServer():Promise<Array<manga.Server>>;

export function InsertManga(arg1:manga.Manga):Promise<number>;

export function TestHomeApiQuery():Promise<any>;

export function UpdateChapJS(arg1:number,arg2:string):Promise<boolean>;

export function UpdatePagesJS(arg1:number,arg2:string):Promise<boolean>;
