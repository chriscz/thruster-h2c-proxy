require_relative "lib/thruster_h2c_proxy/version"

Gem::Specification.new do |s|
  s.name        = "thruster-h2c-proxy"
  s.version     = Thruster::VERSION
  s.summary     = "HTTP/2 proxy with Cleartext"
  s.description = "A fork of Basecamp Thruster but with H2C support"
  s.authors     = [ "Hendre Hayman" ]
  s.email       = "hendrehayman@gmail.com"
  s.homepage    = "https://github.com/skulos/thruster-h2c-proxy"
  s.license     = "MIT"

  s.metadata = {
    "homepage_uri" => s.homepage,
    "rubygems_mfa_required" => "true"
  }

  s.files = Dir[ "{lib}/**/*", "MIT-LICENSE", "README.md" ]
  s.bindir = "exe"
  s.executables << "thruster_h2c_proxy"
end
