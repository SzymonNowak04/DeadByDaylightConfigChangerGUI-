# Config changer for Dead By Deaylight

## For now only for epic store version
And if only installed under C:\ further updates will include support for versions outside epic games store and support for other drives than C:\

### Everything you're changing it's on your own resposibility

### Recommended is that you'll make backup of file in

```text
C:\Users\userName\AppData\Local\DeadByDaylight\Saved\Config\EGS
```

file is named : `GameUserSettings.ini`
(also can be your backup one, but note i have my own settings)
and should look like this:

```ini
[/Script/Engine.GameUserSettings]
bUseDesiredScreenHeight=False

#in this group 0 is minimum value and 3 is maximum
[ScalabilityGroups]
sg.ResolutionQuality=100.000000 #resolution scaling
sg.ViewDistanceQuality=3 # view distance
sg.AntiAliasingQuality=2 # antialiasing
sg.ShadowQuality=3 #shadows
sg.PostProcessQuality=2 # postproccessing
sg.TextureQuality=3 # quality of textures
sg.EffectsQuality=3 # quality of effects
sg.FoliageQuality=3 # quality of foliage (amount of grass, rocks and other small objects on the ground)
sg.ShadingQuality=3 # general shading quality, if you want to see more pixelated effects go ahead and change to 0
sg.AnimationQuality=3 # amount and quality of animations

[/Script/DeadByDaylight.DBDGameUserSettings]
DeviceLoginTokenID= #private stuff
ScreenScaleForWindowedMode=0 
ScalabilityLevel=0 # yes
AutoScalabilitySet=False # auto scaling for best performance
AutoAdjust=False # auto adjusting scaleability
ScreenResolution=100 # yes
FullScreen=True # enable fullscreen
MenuScaleFactor=100 # scale of menu
HudScaleFactor=100 # scale of heads up display
SkillCheckScaleFactor=100 # this wheel when you need to press space scale
LargeText=False # large text for better accessability
HUDPlayerNamesVisibility=True # show players name on HUD
HUDKillerHookCountVisibility=True # HUD for killer how many people are on the hook
HUDScoreEventsVisibility=True # showing *score in events i guess*
FPSLimit=120 # fps limit - DO NOT CHANGE THAT to more than 120 because game is capped at 120 in any case
MainVolume=25 # master volume
MainVolumeOn=True # volume in main menu
MenuMusicVolume=100 # menu music volume
MenuMusicVolumeOn=True # toggle menu music
UseHeadphones=True # headphone mode of sound
MuteOnFocusLost=True # toggle when you want this game to mute when lost focus
KillerCameraSensitivity=50 # killer camera sensivitiy
SurvivorCameraSensitivity=50 # survivor camera sensitivity
KillerMouseSensitivity=50 
SurvivorMouseSensitivity=50
KillerControllerSensitivity=50
SurvivorControllerSensitivity=50
AimAssist=False #toggle aim assist
ControlType=1 # controls layout
InvertY=False 
SurvivorInvertY=False
Language=en # game language
LanguageIsDefinedByPlayer=False
HighestWeightSeenNews=3015
SharedLoginInformation=(LoginProvider="",AuthToken="",TokenType="")
LastPanelContextId=0
ArchivesAutoPlayVoiceOver=True
HasAcceptedCrossplayPopup=True
HasAcceptedCrossProgressionPopup=True
UseAtlantaCustomizedHuds=False
UseAtlantaSurvivorQuickTurn=True
UseAtlantaKillerQuickTurn=True
ShowPortraitBorder=False
PartyPrivacyState=friends
ColorBlindMode=0
ColorBlindModeIntensity=0
BeginnerMode=True
Subtitles=False
SubtitlesBackgroundOpacity=1
SubtitlesSize=1
bUseVSync=True
bUseDynamicResolution=False
ResolutionSizeX=1920
ResolutionSizeY=1080
LastUserConfirmedResolutionSizeX=1920 
LastUserConfirmedResolutionSizeY=1080
WindowPosX=-1 # better do not change that
WindowPosY=-1 # better do not change that
FullscreenMode=1 # fullscreen mode i suppose it means if fulllscreen borderless or no
LastConfirmedFullscreenMode=1 # better do not change that
PreferredFullscreenMode=1  # better do not change that
Version=5 # game version
AudioQualityLevel=2 # 0 minimum; 2 is maximum; changing the audio quality
LastConfirmedAudioQualityLevel=2 # last confirmed level of audio quality
FrameRateLimit=120.000000 # like previous do not change this one above 120 and DO NOT insert negative values
DesiredScreenWidth=1920 # your desired screen width 
DesiredScreenHeight=1080 # your desired screen height
LastUserConfirmedDesiredScreenWidth=1920 # screen width in pixels
LastUserConfirmedDesiredScreenHeight=1080 # screen height in pixels
LastRecommendedScreenWidth=-1.000000 # better do not change
LastRecommendedScreenHeight=-1.000000 # better do not change
LastCPUBenchmarkResult=-1.000000 # better do not change
LastGPUBenchmarkResult=-1.000000 # better do not change
LastGPUBenchmarkMultiplier=1.000000 # better do not change
bUseHDRDisplayOutput=False # toggle HDR
HDRDisplayOutputNits=1000 # changing of HDR display brighntess in nits

[ShaderPipelineCache.CacheFile]
LastOpened=DeadByDaylight # yes
```
