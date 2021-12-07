# frozen_string_literal: true

RSpec.describe 'Backoffice::Categories', type: :request do
  pending do
    it_behaves_like 'edit', '/backoffice/categories/54/edit.js' do
      before { create :category, id: 54 }
    end
  end

  pending do
    it_behaves_like 'update', '/backoffice/categories/54.js' do
      let :valid_params do
        { category: { name: 'Drinks' } }
      end

      let :invalid_params do
        { category: { name: '' } }
      end

      before { create :category, id: 54 }

      let(:success) { -> { should render_template :update } }

      let(:failure) { -> { should render_template :edit } }
    end
  end

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
    before { post '/backoffice/categories', params: params, xhr: true }

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
end
