import {
  createGlobalState,
  StorageSerializers,
  useStorage,
  toReactive,
} from '@vueuse/core'
import { manga } from '@wails/go/models'

/* INTERFACE */
export interface Nav {
  page: number
  limit: number
}

export const useMangaState = createGlobalState(() => {
  // data for homepage mangahome
  const mangaHome = useStorage<manga.MangaHome | null>(
    'mangahome',
    null,
    sessionStorage,
    { serializer: StorageSerializers.object }
  )

  const navStorage = useStorage<Nav>(
    'navhome',
    { page: 1, limit: 10 },
    sessionStorage
  )

  const navHome = toReactive(navStorage)

  const searchManga = useStorage('searchmangahome', '', sessionStorage)

  return { mangaHome, navHome, searchManga }
})
