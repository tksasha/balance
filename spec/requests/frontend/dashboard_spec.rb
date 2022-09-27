# frozen_string_literal: true

RSpec.describe 'Frontend/Dashboard', type: :request do
  describe 'GET /frontend/dashboard.html' do
    before { get '/frontend/dashboard' }

    it_behaves_like 'show.html'
  end
end
