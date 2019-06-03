# iui

Create [IntelliJ](https://www.jetbrains.com/idea/) [ui themes](https://blog.jetbrains.com/idea/2019/03/brighten-up-your-day-add-color-to-intellij-idea/) from either a intellij editor theme or by supplying a json map of theme colors.

# Install

` go get -u github.com/mswift42/iui`

Download ui theme text template from [here](https://github.com/mswift42/iui/blob/master/templ.txt).


# Usage

Create custom ui theme folder by following [these instructions](https://www.jetbrains.org/intellij/sdk/docs/reference_guide/ui_themes/themes.html).

If you have alreadey an IntelliJ Editor Theme in format of <themename>.icls or <themename>.xml you can generate
the ui theme with 

If you do not have an editor theme, you can create one with [themecreator](https://mswift42.github.com/themecreator).

Then you can generate the ui theme with

`iui generate <path to editor theme> <path to ui template>`.


Else you'll have to download the [ThemeColors Json File](https://github.com/mswift42/iui/blob/master/ThemeColors.json) and customize it with your colors,
and generate the ui theme with

`iui generate --json <theme colors>.json <path to ui template>`.

