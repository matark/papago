require 'json'
require 'open-uri'

require 'papago/version'
require 'papago/translator'
require 'papago/configuration'

module Papago
  module_function

  def translate(text)
    Translator.new(text)
  end

  def [](name)
    configuration.send(name)
  end
end
