import {
  createGlobalState,
  StorageSerializers,
  useStorage,
  toReactive,
} from '@vueuse/core'
import { GetMangaHome } from '@wails/go/manga/Manga'
import { GetBreakPoints } from '@/composable/helper'
import { manga } from '@wails/go/models'

/* INTERFACE */
export interface Nav {
  page: number
  per_page: number
}

export const useMangaState = createGlobalState(() => {
  // data for homepage mangahome
  const mangaHome = useStorage<manga.MangaHome | null>(
    'mangahome',
    null,
    sessionStorage,
    { serializer: StorageSerializers.object }
  )

  const dateFilter = useStorage<Date>('date_filter', new Date(), sessionStorage)
  const searchManga = useStorage('searchmangahome', '', sessionStorage)

  const { breakpoints } = GetBreakPoints()
  const lg = breakpoints.greater('lg')

  const getLimit = <T>(min: T, high: T) => {
    return computed<T>(() => {
      if (lg.value) {
        return high
      } else {
        return min
      }
    })
  }

  const useNavStorage = (limit: globalThis.ComputedRef<number>) => {
    const navStorage = useStorage<Nav>(
      'navhomeStorage',
      { page: 1, per_page: 10 },
      sessionStorage
    )

    const navHome = toReactive(navStorage)

    //watch large
    watch(lg, () => (navStorage.value.per_page = limit.value))

    return navHome
  }

  const totalManga = ref(0)
  const loadManga = async () => {
    if (!mangaHome.value) {
      mangaHome.value = await GetMangaHome(null, null, 0, 0)
      totalManga.value = mangaHome.value.manga.length
    }
  }

  return {
    mangaHome,
    searchManga,
    dateFilter,
    getLimit,
    useNavStorage,
    loadManga,
    totalManga,
    lg,
  }
})
