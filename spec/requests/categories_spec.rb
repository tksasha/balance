# frozen_string_literal: true

RSpec.describe 'Categories', type: :request do
  it_behaves_like 'new', uri: '/categories/new.js'
end
