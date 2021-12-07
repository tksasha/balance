# frozen_string_literal: true

RSpec.describe 'Consolidations', type: :request do
  describe 'GET /index.js' do
    before { create_list :item, 2 }

    before { get '/consolidations', xhr: true }

    it_behaves_like 'index.js'
  end
end
