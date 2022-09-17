# frozen_string_literal: true

RSpec.describe 'Backoffice::Cashes', type: :request do
  describe 'GET /index.js' do
    before do
      create_list :cash, 2

      get '/backoffice/cashes', xhr: true
    end

    it_behaves_like 'index.js'
  end

  describe 'GET /new.js' do
    before { get '/backoffice/cashes/new', xhr: true }

    it_behaves_like 'new.js'
  end

  describe 'POST /create.js' do
    before { post '/backoffice/cashes', params:, xhr: true }

    context 'with valid params' do
      let(:params) { { cash: { name: 'Bank', formula: '4 + 5', supercategory: 'bonds' } } }

      let(:cash) { Cash.last }

      it_behaves_like 'create.js'

      it { expect(cash.supercategory).to eq 'bonds' }
    end

    context 'with invalid params' do
      let(:params) { { cash: { name: '' } } }

      it_behaves_like 'new.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end

  describe 'GET /edit.js' do
    let(:cash) { create :cash }

    before { get "/backoffice/cashes/#{ cash.id }/edit", xhr: true }

    it_behaves_like 'edit.js'
  end

  describe 'PATCH /update.js' do
    let(:cash) { create :cash }

    before { patch "/backoffice/cashes/#{ cash.id }", params:, xhr: true }

    context 'with valid params' do
      before { cash.reload }

      let(:params) { { cash: { name: Faker::Lorem.word, supercategory: 'bonds' } } }

      it_behaves_like 'update.js'

      it { expect(cash.supercategory).to eq 'bonds' }
    end

    context 'with invalid params' do
      let(:params) { { cash: { name: '' } } }

      it_behaves_like 'edit.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end

  describe 'DELETE /destroy.js' do
    let(:cash) { create :cash }

    before { delete "/backoffice/cashes/#{ cash.id }", xhr: true }

    it_behaves_like 'destroy.js'

    it { expect { cash.reload }.to raise_error(ActiveRecord::RecordNotFound) }
  end
end
