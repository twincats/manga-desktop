import { EventsOnce, EventsOn, EventsEmit } from '@wails/runtime/runtime'

export const EmitListenOnce = (
  eventName: string,
  callback: (...data: any) => void,
  ...extraData: any
) => {
  EventsEmit(eventName, extraData)
  EventsOnce(eventName + '_server', callback)
}

export const Listen = EventsOn
export const ListenOnce = EventsOnce
export const Emit = EventsEmit
