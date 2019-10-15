# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  it_behaves_like 'new', uri: '/cashes/new.js'
end
