# frozen_string_literal: true

source 'https://rubygems.org'

ruby '~> 3.3'

gem 'activeadmin'
gem 'annotate'
gem 'bootsnap', require: false
gem 'bootstrap'
gem 'draper'
gem 'faker'
gem 'jquery-rails'
gem 'mini_racer'
gem 'month'
gem 'paper_trail'
gem 'puma'
gem 'rails'
gem 'rake'
gem 'sassc-rails'
# https://github.com/rails/sprockets-rails/issues/524
gem 'sprockets-rails', '3.4.2'
# https://github.com/rails/rails/blob/747a03ba7722b6f0a7ce42e86cea83cf07a2e6ef/activerecord/lib/active_record/connection_adapters/sqlite3_adapter.rb#L14
gem 'sqlite3', '~> 1.4' # TODO: remove it after updating ActiveRecord
gem 'uglifier'

group :development do
  gem 'brakeman', require: false
  gem 'bundler-audit', require: false
  gem 'bundler-leak', require: false
  gem 'fasterer', require: false
  gem 'listen'
  gem 'rubocop', '1.63.1', require: false # TODO: there is a problem with 1.63.2
  gem 'rubocop-factory_bot', require: false
  gem 'rubocop-performance', require: false
  gem 'rubocop-rails', require: false
  gem 'rubocop-rake', require: false
  gem 'rubocop-rspec', require: false
  gem 'rubocop-rspec_rails', require: false
end

group :development, :test do
  gem 'bullet'
  gem 'debug'
  gem 'rspec-rails'
end

group :test do
  gem 'database_cleaner'
  gem 'factory_bot_rails'
  gem 'rails-controller-testing'
  gem 'rspec-activemodel-mocks'
  gem 'rspec-its'
  gem 'shoulda-callback-matchers'
  gem 'shoulda-matchers'
end
