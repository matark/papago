Gem::Specification.new do |spec|
  spec.name       = 'papago'
  spec.version    = '1.0.1'
  spec.license    = 'MIT'
  spec.author     = 'minisode'
  spec.email      = 'minisode@189.cn'
  spec.executable = 'papago'
  spec.files      = Dir['lib/**/*']
  spec.summary    = 'A simple command-line translator'
  spec.homepage   = 'https://minisode.surge.sh/papago'

  spec.add_development_dependency 'minitest', '~> 5.11'
  spec.add_runtime_dependency 'activesupport', '~> 5.2'
  spec.add_runtime_dependency 'mercenary', '~> 0.3.6'
  spec.add_runtime_dependency 'lolcat', '~> 99.9'
  spec.add_runtime_dependency 'treely', '~> 1.0'
end
