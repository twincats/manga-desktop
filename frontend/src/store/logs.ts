import { createGlobalState, useStorage } from '@vueuse/core'
import { DateApp } from '@/composable/helper'
import { EventChap, EventPage } from '@/type/download'

export const useLogState = createGlobalState(() => {
  const log_dl = useStorage<string[]>('log_download', [], sessionStorage)

  const log = (
    message: EventPage | EventChap | string,
    stat: 'Fail' | 'Danger' | 'Success' | 'Warning' | 'other' = 'other'
  ) => {
    let loginfo = get_statusLog(stat)
    let message_data = ''
    if (message instanceof EventChap) {
      loginfo = get_statusLog('ChapterList')
      message_data = `[Chapter ${message.chapter}] Total Pages : ${message.total_page}`
    } else if (message instanceof EventPage) {
      if (message.stat_error) {
        loginfo = get_statusLog('Fail')
        message_data = `[Chapter ${message.chapter}] [Image URL] : ${message.images}`
      } else {
        loginfo = get_statusLog('Success')
        message_data = `[Chapter ${message.chapter}] [Image URL] : ${message.images}`
      }
    } else {
      message_data = message
    }
    //add log
    log_dl.value.push(
      `${loginfo} [${DateApp.NewDate().Format(
        'DD-MM-YYYY HH:mm:ss.ms'
      )}] : ${message_data}`
    )
  }

  const get_statusLog = (stat: string): string => {
    let loginfo = ''
    switch (stat) {
      case 'Fail':
        loginfo = "<span class='logfail'>[Fail]</span>"
        break
      case 'Danger':
        loginfo = "<span class='logdanger'>[Error]</span>"
        break
      case 'Warning':
        loginfo = "<span class='logwarning'>[Warning]</span>"
        break
      case 'Success':
        loginfo = "<span class='logsuccess'>[Success]</span>"
        break
      case 'other':
        loginfo = "<span class='logother'>[Other]</span>"
        break
      default:
        loginfo = `<span>[${stat}]</span>`
    }
    return loginfo
  }

  const get_logtext = () => {
    const resultLogRaw = log_dl.value.join('\n')
    var doc = new DOMParser().parseFromString(resultLogRaw, 'text/html')
    return doc.body.textContent || ''
  }

  const log_clear = () => {
    log_dl.value = []
  }
  return {
    log_dl,
    log,
    log_clear,
    get_logtext,
  }
})
