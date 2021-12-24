# frozen_string_literal: true

RSpec.describe 'Backoffice::Dashboard', type: :request do
  describe 'GET /show.js' do
    before { get '/backoffice', xhr: true }

    it_behaves_like 'show.js'
  end
end
