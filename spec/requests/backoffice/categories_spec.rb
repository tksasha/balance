# frozen_string_literal: true

RSpec.describe 'Backoffice::Categories', type: :request do
  describe 'GET /index.js' do
    before { create_list :category, 2 }

    before { get '/backoffice/categories', xhr: true }

    it_behaves_like 'index.js'
  end

  describe 'GET /new.js' do
    before { get '/backoffice/categories/new', xhr: true }

    it_behaves_like 'new.js'
  end

  describe 'POST /create.js' do
    before { post '/backoffice/categories', params:, xhr: true }

    context 'with valid params' do
      let(:params) { { category: { name: 'Drinks', income: true, visible: true } } }

      it_behaves_like 'create.js'
    end

    context 'with invalid params' do
      let(:params) { { category: { name: '' } } }

      it_behaves_like 'new.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end

  describe 'GET /edit.js' do
    let(:category) { create :category }

    before { get "/backoffice/categories/#{ category.id }/edit", xhr: true }

    it_behaves_like 'edit.js'
  end

  describe 'PATCH /update.js' do
    let(:category) { create :category }

    before { patch "/backoffice/categories/#{ category.id }", params:, xhr: true }

    context do
      let(:params) { { category: { name: 'Drinks' } } }

      it_behaves_like 'update.js'
    end

    context do
      let(:params) { { category: { name: '' } } }

      it_behaves_like 'edit.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end
end
