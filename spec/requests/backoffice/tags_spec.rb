# frozen_string_literal: true

RSpec.describe 'Backoffice::Tags', type: :request do
  describe 'GET /index.js' do
    let(:category) { create :category }

    before { create_list :tag, 2, category: category }

    before { get "/backoffice/categories/#{ category.id }/tags", xhr: true }

    it_behaves_like 'index.js'
  end

  describe 'GET /new.js' do
    let(:category) { create :category }

    before { get "/backoffice/categories/#{ category.id }/tags/new", xhr: true }

    it_behaves_like 'new.js'
  end

  describe 'POST /create.js' do
    let(:category) { create :category }

    before { post "/backoffice/categories/#{ category.id }/tags", params: params, xhr: true }

    context 'with valid params' do
      let(:params) { { tag: { name: Faker::Commerce.color } } }

      it_behaves_like 'create.js'
    end

    context 'with invalid params' do
      let(:params) { { tag: { name: '' } } }

      it_behaves_like 'new.js'
    end
  end
end
