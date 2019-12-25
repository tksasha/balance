# frozen_string_literal: true

RSpec.describe 'Backoffice::ExchangeRates', type: :request do
  it_behaves_like 'index', '/backoffice/exchange_rates.js'
end
