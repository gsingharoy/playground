require "http/server"

server = HTTP::Server.new(3100) do |context|
  context.response.content_type = "text/plain"
  context.response.print "Hello World again !! Its now #{Time.now}"
end

puts "Listening on http://0.0.0.0:3100"
server.listen
