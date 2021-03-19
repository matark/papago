module Papago
  class Configuration
    attr_accessor :key
    attr_accessor :keyfrom

    def api_key=(api_key)
      @keyfrom, @key = api_key&.gsub(/[\w]+/).to_a
    end
  end

  def self.configuration
    @configuration ||= Configuration.new
  end

  def self.configuration=(config)
    @configuration = config
  end

  def self.configure
    yield configuration
  end
end
