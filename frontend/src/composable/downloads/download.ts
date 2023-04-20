//download composable to make downloads page more compact
import { GetBreakPoints } from '@/composable/helper'
import { manga, types } from '@wails/go/models'
import { GetServer } from '@wails/go/manga/Manga'

export const UseTable = () => {
  const tableDownload = reactive<types.ChapterList[]>([]) // data for display chapter in table
  const tablePageSize = ref(5) // data shown how much data perpage in table default 5=>15
  const tableLoading = ref(false) // data for status loading table
  const tablePageCurrent = ref(1)
  const { lg } = GetBreakPoints() // global breakpoint
  // set default perpage table
  if (lg.value) {
    tablePageSize.value = 5
  } else {
    tablePageSize.value = 15
  }

  // watch perpage for dynamic change when windows size change
  watch(lg, nlg => {
    if (nlg) {
      tablePageSize.value = 5
    } else {
      tablePageSize.value = 15
    }
  })

  return {
    tableDownload,
    tablePageSize,
    tableLoading,
    tablePageCurrent,
  }
}

export const UseServer = () => {
  const servers = ref<manga.Server[] | null>(null) // to fill list server
  GetServer().then(res => {
    // to fetch server only active server
    servers.value = res.filter(item => item.status_active === true)
  })
  const getSelectedServer = (id: number): manga.Server | undefined => {
    return servers.value?.find(item => item.id == id)
  }
  return {
    servers,
    getSelectedServer,
  }
}
