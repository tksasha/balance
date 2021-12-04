# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  describe 'GET /index.js' do
    before { create_list :cash, 2 }

    before { get '/cashes', xhr: true }

    it_behaves_like 'index.js'
  end

  describe 'GET /edit.js' do
    let(:cash) { create :cash }

    before { get "/cashes/#{ cash.id }/edit", xhr: true }

    it_behaves_like 'edit.js'
  end

  describe 'PATCH /update.js' do
    let(:cash) { create :cash }

    before { patch "/cashes/#{ cash.id }", params: params, xhr: true }

    context 'with valid params' do
      let(:params) { { cash: { name: Faker::Lorem.word } } }

      it_behaves_like 'update.js'
    end

    context 'with invalid params' do
      let(:params) { { cash: { name: '' } } }

      it_behaves_like 'edit.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end
end
