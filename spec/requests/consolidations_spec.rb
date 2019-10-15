# frozen_string_literal: true

RSpec.describe 'Consolidations', type: :request do
  it_behaves_like 'index', uri: '/consolidations.js'
end
