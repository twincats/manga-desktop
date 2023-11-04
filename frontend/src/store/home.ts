import {
  createGlobalState,
  StorageSerializers,
  useStorage,
  toReactive,
  isDefined,
} from '@vueuse/core'
import { GetMangaHome } from '@wails/go/manga/Manga'
import { GetBreakPoints } from '@/composable/helper'
import { manga } from '@wails/go/models'
import { GetCountMangas } from '@wails/go/manga/Manga'

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
    localStorage,
    { serializer: StorageSerializers.object }
  )

  //check count if mangaHome available (clear mangaHome when not match DB)
  if (isDefined(mangaHome)) {
    const mh_length = mangaHome.value.manga.length
    GetCountMangas().then(totalManga => {
      if (totalManga != mh_length) clearManga()
    })
  }

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

    // watch large
    watch(lg, () => (navStorage.value.per_page = limit.value))

    // watch perPage
    let oldPage = navStorage.value.page
    watch(
      () => navStorage.value.per_page,
      (npp, opp) => {
        if (npp > opp) oldPage = navStorage.value.page
        const np = Math.floor(((navStorage.value.page - 1) * opp) / npp) + 1
        const npo = Math.floor(((oldPage - 1) * npp) / opp) + 1
        if (npp < opp && oldPage > np && np >= npo) {
          navStorage.value.page = oldPage
        } else {
          navStorage.value.page = np > 0 ? np : 1
        }
      }
    )

    return navHome
  }

  const totalManga = ref(0)
  const loading = ref(false)
  const loadManga = async () => {
    if (!isDefined(mangaHome)) {
      loading.value = true
      mangaHome.value = await GetMangaHome(null, null, 0, 0)
      totalManga.value = mangaHome.value.manga.length
      loading.value = false
    }
  }

  const clearManga = () => {
    mangaHome.value = null
  }

  return {
    mangaHome,
    searchManga,
    dateFilter,
    getLimit,
    useNavStorage,
    loading,
    loadManga,
    totalManga,
    clearManga,
    lg,
  }
})
