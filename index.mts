import md, { type Link, type Text } from 'markdown-ast'
import { type DefaultTheme } from 'vitepress'
import { Project, SyntaxKind, ts } from "ts-morph"

const summaryFile = Bun.file('./docs/summary.md')
const vitepressConfigFilePath = './.vitepress/config.ts'

async function saveConfigFile(sidebarData: DefaultTheme.SidebarItem[]) {
  // initialize
  const project = new Project({
    tsConfigFilePath: './tsconfig.json'
  })
  const file = project.addSourceFileAtPath(vitepressConfigFilePath)
  const exportAssignment = file.getExportAssignmentOrThrow(exportAssignment => true)
  const objectLiteral = exportAssignment.getFirstDescendantByKindOrThrow(SyntaxKind.ObjectLiteralExpression)
  const themeConfig = objectLiteral.getPropertyOrThrow('themeConfig')
  const properties = themeConfig.getFirstDescendantByKind(SyntaxKind.ObjectLiteralExpression)
  const sidebar = properties!.getPropertyOrThrow('sidebar')
  const elements = sidebar.getFirstDescendantByKind(SyntaxKind.ArrayLiteralExpression)
  // clear
  if (elements) {
    for (let i = 0; i < elements.getElements().length; i++) {
      elements.removeElement(i)
    }
  }
  // add sidebar data
  // elements?.addElements(sidebarData.map(createSidebarItem))
  elements?.addElements(sidebarData.map(item => JSON.stringify(item)))
  file.formatText({
    semicolons: ts.server.protocol.SemicolonPreference.Ignore,
    convertTabsToSpaces: true,
    ensureNewLineAtEndOfFile: true,
    indentStyle: ts.IndentStyle.Smart,
    indentSize: 2,
    placeOpenBraceOnNewLineForFunctions: true,
    tabSize: 2,
    trimTrailingWhitespace: true,
    insertSpaceAfterCommaDelimiter: true,   
    indentMultiLineObjectLiteralBeginningOnBlankLine: true,
    insertSpaceAfterOpeningAndBeforeClosingJsxExpressionBraces: true,
    newLineCharacter: '\n' 
  })

  // save
  await project.save()
}

async function run() {
  let summary = await summaryFile.text()
  // replace for specific scene
  summary = summary.replace(/<(.+\.md)>/g, '$1')
  const ast = md(summary)
  const sidebarData: DefaultTheme.SidebarItem[] = []
  const treeStack: DefaultTheme.SidebarItem[] = []
  let prevLevel = 0
  for (const line of ast) {
    // title
    if (line.type === 'title') {
      // toc
      if (line.rank === 1) {
        //
      } else if (line.rank === 2) {
        //
      }

    } else if (line.type === 'list') {
      const level = line.indent.length / 2
      const link = line.block[0] as Link
      const block = link.block[0] as Text
      const url = link.url.replace(/\.md$/, '')

      const item: DefaultTheme.SidebarItem = {
        text: block.text,
        link: url === 'README' ? '/' : url
      }
      const prev = treeStack[treeStack.length - 1]
      if (level > prevLevel) {
        if (!prev.items) {
          prev.items = []
        }
        if (prevLevel === 0) {
          prev.collapsed = true
        } else {
          prev.collapsed = false
        }
        treeStack.push(item)
        prev.items.push(item)
      } else if (level === prevLevel) {
        sidebarData.push(item)
        if (!prev) {
          treeStack.push(item)
        } else {
          treeStack[treeStack.length - 1] = item
        }
      } else {
        for (let i = 0; i < prevLevel - level; i++) {
          treeStack.pop()
        }
        sidebarData.push(item)
      }
      prevLevel = level
    }
  }

  saveConfigFile(sidebarData)
}

run()