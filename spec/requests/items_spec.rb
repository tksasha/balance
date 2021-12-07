# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
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
            description: 'Lorem Ipsum ...'
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

  describe 'GET /edit.js' do
    let(:item) { create :item }

    before { get "/items/#{ item.id }/edit", xhr: true }

    it_behaves_like 'edit.js'
  end

  describe 'PATCH /update.js' do
    let(:item) { create :item }

    before { patch "/items/#{ item.id }", params: params, xhr: true }

    context 'with valid params' do
      let(:params) { { item: { date: '2019-11-13' } } }

      it_behaves_like 'update.js'
    end

    context 'with invalid params' do
      let(:params) { { item: { date: '' } } }

      it_behaves_like 'edit.js'

      it { expect(response).to have_http_status(:unprocessable_entity) }
    end
  end

  describe 'DELETE /destroy.js' do
    let(:item) { create :item }

    before { delete "/items/#{ item.id }", xhr: true }

    it_behaves_like 'destroy.js'

    it { expect(item.reload.deleted_at).not_to be_nil }
  end
end
