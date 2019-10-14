# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  it_behaves_like 'index', uri: '/items'

  it_behaves_like 'index', uri: '/items.js'
end
