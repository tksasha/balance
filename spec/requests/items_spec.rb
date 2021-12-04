# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  pending do
    it_behaves_like 'destroy', '/items/47.js' do
      before { create :item, id: 47 }

      let(:success) { -> { should render_template :destroy } }
    end
  end

  pending do
    it_behaves_like 'edit', '/items/47/edit.js' do
      before { create :item, id: 47 }
    end
  end

  pending do
    it_behaves_like 'update', '/items/47.js' do
      let :valid_params do
        { item: { date: '2019-11-13' } }
      end

      let :invalid_params do
        { item: { date: '' } }
      end

      before { create :item, id: 47 }

      let(:success) { -> { should render_template :update } }

      let(:failure) { -> { should render_template :edit } }
    end
  end

  describe 'GET /index' do
    before { get '/items' }

    it_behaves_like 'index.html'
  end

  describe 'GET /index.js' do
    before { create_list :item, 2 }

    before { get '/items', xhr: true }

    it_behaves_like 'index.js'
  end

  describe 'POST /create.js' do
    let(:category) { create :category }

    before { post '/items', params: params, xhr: true }

    context 'with valid params' do
      let(:params) do
        {
          item: {
            date: '2019-11-13',
            formula: '3 + 5',
            category_id: category.id,
            description: 'Lorem Ipsum ...',
          }
        }
      end

      it_behaves_like 'create.js'
    end

    context 'with invalid params' do
      let(:params) { { item: { date: '' } } }

      it_behaves_like 'new.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end
end
