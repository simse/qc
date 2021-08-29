const path = require("path")

exports.createPages = async ({ graphql, actions, reporter }) => {
    const { createPage } = actions
    // Query for markdown nodes to use in creating pages.
    const result = await graphql(
        `
      {
        allMarkdownRemark(filter: {fileAbsolutePath: {regex: "/docs/gm"}}) {
          nodes {
            id
            frontmatter {
              title
            }
            fileAbsolutePath
          }
        }
      }
      `
    )

    // Handle errors
    if (result.errors) {
        reporter.panicOnBuild(`Error while running GraphQL query.`)
        return
    }

    //console.log(result)

    // Create pages for each markdown file.
    const docsPageTemplate = path.resolve(`src/templates/docs-page.js`)
    result.data.allMarkdownRemark.nodes.forEach(( node ) => {
        // Generate path
        let string = node.fileAbsolutePath
        let split_string = string.split("/docs/")
        let filePath = split_string[split_string.length - 1].replace(".md", "")
        let parts = filePath.split("/")
        let path = "/docs/" + parts.filter(part => part !== "index").join("/") 

        createPage({
            path,
            component: docsPageTemplate,
            context: {
                id: node.id
            },
        })
    })
}