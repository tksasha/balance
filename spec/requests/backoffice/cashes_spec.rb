# frozen_string_literal: true

RSpec.describe 'Backoffice::Cashes', type: :request do
  pending do
    it_behaves_like 'update', '/backoffice/cashes/49.js' do
      let :valid_params do
        { cash: { name: 'Bank' } }
      end

      let :invalid_params do
        { cash: { name: '' } }
      end

      before { create :cash, id: 49 }

      let(:success) { -> { should render_template :update } }

      let(:failure) { -> { should render_template :edit } }
    end
  end

  pending do
    it_behaves_like 'destroy', '/backoffice/cashes/49.js' do
      before { create :cash, id: 49 }

      let(:success) { -> { should render_template :destroy } }
    end
  end

  describe 'GET /index.js' do
    before { create_list :cash, 2 }

    before { get '/backoffice/cashes', xhr: true }

    it_behaves_like 'index.js'
  end

  describe 'GET /new.js' do
    before { get '/backoffice/cashes/new', xhr: true }

    it_behaves_like 'new.js'
  end

  describe 'POST /create.js' do
    before { post '/backoffice/cashes', params: params, xhr: true }

    context 'with valid params' do
      let(:params) { { cash: { name: 'Bank', formula: '4 + 5' } } }

      it_behaves_like 'create.js'
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
end
