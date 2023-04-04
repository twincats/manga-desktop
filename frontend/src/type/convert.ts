export interface StartConvertEventData {
  size_before: number
  total_convert: number
}

export interface RunConvertEventData {
  error: string
  filepath: string
  filename: string
  chap: string
  resized: boolean
  deleted: boolean
}
