module Papago
  class Translator
    def initialize(text)
      @request = URI('https://fanyi.youdao.com/openapi.do')
      @request.query = URI.encode_www_form({
        :keyfrom => Papago[:keyfrom],
        :key     => Papago[:key],
        :type    => 'data',
        :doctype => 'json',
        :version => '1.1',
        :q       => text
      })

      build_context!
    end

    def contents
      [
        {
          name: "#{@context.fetch(:query)} [#{@context.dig(:basic, :phonetic)}]",
          contents: [
            {
              name: @context.fetch(:translation).join(%{ }),
              contents: (@context.dig(:basic, :explains) || []).map { |explain| { name: explain } }
            }
          ] + @context.fetch(:web, []).map { |explain| { name: "#{explain.fetch(:key)}: #{explain.fetch(:value).join(%{ })}" } }
        }
      ]
    end

    def build_context!
      @context = JSON.parse(@request.read, symbolize_names: true)
    end
  end
end
