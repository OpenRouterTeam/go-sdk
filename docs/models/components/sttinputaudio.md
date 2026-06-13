# STTInputAudio

Base64-encoded audio to transcribe


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Data`                                                                                        | `string`                                                                                      | :heavy_check_mark:                                                                            | Base64-encoded audio data (raw bytes, not a data URI)                                         |
| `Format`                                                                                      | `string`                                                                                      | :heavy_check_mark:                                                                            | Audio format (e.g., wav, mp3, flac, m4a, ogg, webm, aac). Supported formats vary by provider. |