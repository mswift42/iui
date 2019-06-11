# iui

Create [IntelliJ](https://www.jetbrains.com/idea/) [ui themes](https://blog.jetbrains.com/idea/2019/03/brighten-up-your-day-add-color-to-intellij-idea/) from either a intellij editor theme or by supplying a json map of theme colors.

# Install

Go to [releases](https://github.com/mswift42/iui/releases), download your zip archive and extract it.

# Usage

Create custom ui theme folder by following [these instructions](https://www.jetbrains.org/intellij/sdk/docs/reference_guide/ui_themes/themes.html).


If you do not have an editor theme, you can create one with [themecreator](https://mswift42.github.com/themecreator).

Generate the ui theme with

`iui generate <path to editor theme> templ.txt`.


Else you will have to customize the ThemeColors.json file with your colors, and generate the ui theme with

`iui generate --json ThemeColors.json templ.txt`.


Replace the <theme name>.theme.json file in custom ui theme folder with your generated file, make sure to fill
out the author field, press the run button and admire your new theme.

