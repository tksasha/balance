# frozen_string_literal: true

RSpec.describe 'Balances', type: :request do
  it_behaves_like 'show', '/balance.js'
end
