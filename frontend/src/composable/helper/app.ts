import { EventsOnce, EventsOn, EventsEmit } from '@wails/runtime/runtime'

export const EmitListen = (
  eventName: string,
  callback: (...data: any) => void
) => {
  EventsEmit(eventName)
  EventsOn(eventName, callback)
}

export const EmitListenOnce = (
  eventName: string,
  callback: (...data: any) => void
) => {
  EventsEmit(eventName)
  EventsOnce(eventName, callback)
}
