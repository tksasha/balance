# frozen_string_literal: true

RSpec.describe 'Consolidations', type: :request do
  describe 'GET /index.js' do
    before do
      create_list :item, 2

      get '/consolidations', xhr: true
    end

    it_behaves_like 'index.js'
  end
end
