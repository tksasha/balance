# frozen_string_literal: true

RSpec.describe 'Backoffice::Dashboard' do
  describe 'GET /show.js' do
    before { get '/backoffice', xhr: true }

    it_behaves_like 'show.js'
  end
end
