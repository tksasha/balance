require 'rails_helper'

RSpec.describe 'Balances', type: :request do
  it_behaves_like 'show', '/balance.js'
end
