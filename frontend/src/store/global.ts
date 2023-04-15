import { createGlobalState, StorageSerializers } from '@vueuse/core'
import { types } from '@wails/go/models'

export const useDownloadState = createGlobalState(() => {
  // data for table chapter
  const chapterList = useStorage<types.Chapter | null>(
    'my-download-chapter',
    null,
    sessionStorage,
    { serializer: StorageSerializers.object }
  )

  // data for download url chapter
  const urldata = useStorage('my-download-url', '', sessionStorage)

  // data for selected server ID
  const selectedServer = useStorage('my-download-serverID', 1, sessionStorage)

  // data for list selected chapter to download
  const selected_chapter_url = useStorage<types.ChapterList[]>(
    'my-download-selected-dl-chap',
    [],
    sessionStorage
  )
  return { chapterList, urldata, selectedServer, selected_chapter_url }
})
