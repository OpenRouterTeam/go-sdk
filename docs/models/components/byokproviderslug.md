# BYOKProviderSlug

The upstream provider this credential authenticates against, as a lowercase slug (e.g. `openai`, `anthropic`, `amazon-bedrock`).

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.BYOKProviderSlugAi21

// Open enum: custom values can be created with a direct type cast
custom := components.BYOKProviderSlug("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `BYOKProviderSlugAi21`            | ai21                              |
| `BYOKProviderSlugAionLabs`        | aion-labs                         |
| `BYOKProviderSlugAkashml`         | akashml                           |
| `BYOKProviderSlugAlibaba`         | alibaba                           |
| `BYOKProviderSlugAmazonBedrock`   | amazon-bedrock                    |
| `BYOKProviderSlugAmazonNova`      | amazon-nova                       |
| `BYOKProviderSlugAmbient`         | ambient                           |
| `BYOKProviderSlugAnthropic`       | anthropic                         |
| `BYOKProviderSlugArceeAi`         | arcee-ai                          |
| `BYOKProviderSlugAtlasCloud`      | atlas-cloud                       |
| `BYOKProviderSlugAvian`           | avian                             |
| `BYOKProviderSlugAzure`           | azure                             |
| `BYOKProviderSlugBaidu`           | baidu                             |
| `BYOKProviderSlugBaseten`         | baseten                           |
| `BYOKProviderSlugBlackForestLabs` | black-forest-labs                 |
| `BYOKProviderSlugByteplus`        | byteplus                          |
| `BYOKProviderSlugCerebras`        | cerebras                          |
| `BYOKProviderSlugChutes`          | chutes                            |
| `BYOKProviderSlugCirrascale`      | cirrascale                        |
| `BYOKProviderSlugClarifai`        | clarifai                          |
| `BYOKProviderSlugCloudflare`      | cloudflare                        |
| `BYOKProviderSlugCohere`          | cohere                            |
| `BYOKProviderSlugCrusoe`          | crusoe                            |
| `BYOKProviderSlugDarkbloom`       | darkbloom                         |
| `BYOKProviderSlugDecart`          | decart                            |
| `BYOKProviderSlugDeepinfra`       | deepinfra                         |
| `BYOKProviderSlugDeepseek`        | deepseek                          |
| `BYOKProviderSlugDekallm`         | dekallm                           |
| `BYOKProviderSlugDigitalocean`    | digitalocean                      |
| `BYOKProviderSlugFeatherless`     | featherless                       |
| `BYOKProviderSlugFireworks`       | fireworks                         |
| `BYOKProviderSlugFriendli`        | friendli                          |
| `BYOKProviderSlugGmicloud`        | gmicloud                          |
| `BYOKProviderSlugGoogleAiStudio`  | google-ai-studio                  |
| `BYOKProviderSlugGoogleVertex`    | google-vertex                     |
| `BYOKProviderSlugGroq`            | groq                              |
| `BYOKProviderSlugHeygen`          | heygen                            |
| `BYOKProviderSlugInception`       | inception                         |
| `BYOKProviderSlugInceptron`       | inceptron                         |
| `BYOKProviderSlugInferactVllm`    | inferact-vllm                     |
| `BYOKProviderSlugInferenceNet`    | inference-net                     |
| `BYOKProviderSlugInfermatic`      | infermatic                        |
| `BYOKProviderSlugInflection`      | inflection                        |
| `BYOKProviderSlugIoNet`           | io-net                            |
| `BYOKProviderSlugIonstream`       | ionstream                         |
| `BYOKProviderSlugLiquid`          | liquid                            |
| `BYOKProviderSlugMancer`          | mancer                            |
| `BYOKProviderSlugMara`            | mara                              |
| `BYOKProviderSlugMinimax`         | minimax                           |
| `BYOKProviderSlugMistral`         | mistral                           |
| `BYOKProviderSlugModelrun`        | modelrun                          |
| `BYOKProviderSlugModular`         | modular                           |
| `BYOKProviderSlugMoonshotai`      | moonshotai                        |
| `BYOKProviderSlugMorph`           | morph                             |
| `BYOKProviderSlugNcompass`        | ncompass                          |
| `BYOKProviderSlugNebius`          | nebius                            |
| `BYOKProviderSlugNexAgi`          | nex-agi                           |
| `BYOKProviderSlugNextbit`         | nextbit                           |
| `BYOKProviderSlugNovita`          | novita                            |
| `BYOKProviderSlugNvidia`          | nvidia                            |
| `BYOKProviderSlugOpenInference`   | open-inference                    |
| `BYOKProviderSlugOpenai`          | openai                            |
| `BYOKProviderSlugParasail`        | parasail                          |
| `BYOKProviderSlugPerceptron`      | perceptron                        |
| `BYOKProviderSlugPerplexity`      | perplexity                        |
| `BYOKProviderSlugPhala`           | phala                             |
| `BYOKProviderSlugPoolside`        | poolside                          |
| `BYOKProviderSlugRecraft`         | recraft                           |
| `BYOKProviderSlugReka`            | reka                              |
| `BYOKProviderSlugRelace`          | relace                            |
| `BYOKProviderSlugSakanaAi`        | sakana-ai                         |
| `BYOKProviderSlugSambanova`       | sambanova                         |
| `BYOKProviderSlugSeed`            | seed                              |
| `BYOKProviderSlugSiliconflow`     | siliconflow                       |
| `BYOKProviderSlugSourceful`       | sourceful                         |
| `BYOKProviderSlugStepfun`         | stepfun                           |
| `BYOKProviderSlugStreamlake`      | streamlake                        |
| `BYOKProviderSlugSwitchpoint`     | switchpoint                       |
| `BYOKProviderSlugTenstorrent`     | tenstorrent                       |
| `BYOKProviderSlugTogether`        | together                          |
| `BYOKProviderSlugUpstage`         | upstage                           |
| `BYOKProviderSlugVenice`          | venice                            |
| `BYOKProviderSlugWafer`           | wafer                             |
| `BYOKProviderSlugWandb`           | wandb                             |
| `BYOKProviderSlugXai`             | xai                               |
| `BYOKProviderSlugXiaomi`          | xiaomi                            |
| `BYOKProviderSlugZAi`             | z-ai                              |