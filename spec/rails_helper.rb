# frozen_string_literal: true

ENV['RAILS_ENV'] ||= 'test'
require File.expand_path('../config/environment', __dir__)
require 'rspec/rails'
require 'paper_trail/frameworks/rspec'

Dir[Rails.root.join('spec', 'support', '**', '*.rb')].each { |f| require f }

RSpec.configure do |config|
  config.mock_with :rspec

  config.order = :random

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end

  config.include Permitter

  %i[controller view request].each do |type|
    config.include(Rails::Controller::Testing::TestProcess, type:)
    config.include(Rails::Controller::Testing::TemplateAssertions, type:)
    config.include(Rails::Controller::Testing::Integration, type:)
  end

  config.include ActiveSupport::Testing::TimeHelpers

  config.infer_spec_type_from_file_location!
end

Shoulda::Matchers.configure do |config|
  config.integrate do |with|
    with.test_framework :rspec

    with.library :rails
  end
end

# Checks for pending migrations before tests are run.
# If you are not using ActiveRecord, you can remove this line.
ActiveRecord::Migration.maintain_test_schema!
