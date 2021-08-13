const path = require('path')
const client = require('node-grpc-client')
const PROTO_PATH = path.resolve(__dirname, '../blogpb/blog.proto')

const myClient = new client(
  PROTO_PATH,
  'blog',
  'BlogService',
  'localhost:50051'
)

const data = {
  blog: {
    author_id: 'rad',
    title: 'jengjet',
    content: 'Radrad',
  },
}

// myClient.runService('CreateBlog', data, (err, res) => {
//   console.log('Service Response', res)
// })

const options = {
  metadata: {
    hello: 'world',
  },
}
const stream = myClient.listBlogStream({}, options)
stream.on('data', (data) => console.log(data))
console.log(myClient.listNameMethods)
