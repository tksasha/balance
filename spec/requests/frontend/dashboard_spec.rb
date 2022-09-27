# frozen_string_literal: true

RSpec.describe 'Frontend/Dashboard', type: :request do
  describe 'GET /frontend/dashboard.html' do
    before { get '/frontend/dashboard' }

    it_behaves_like 'show.html'
  end

  describe 'GET /frontend/dashboard.js' do
    before { get '/frontend/dashboard', xhr: true }

    it_behaves_like 'show.js'
  end
end
