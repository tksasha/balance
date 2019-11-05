# frozen_string_literal: true

RSpec.describe 'Currencies', type: :request do
  it_behaves_like 'index', '/currencies.js'
end
