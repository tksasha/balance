# frozen_string_literal: true

RSpec.describe 'Backoffice::Tags', type: :request do
  describe 'GET /index.js' do
    let(:category) { create :category }

    before { create_list :tag, 2, category: }

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

    before { post "/backoffice/categories/#{ category.id }/tags", params:, xhr: true }

    context 'with valid params' do
      let(:params) { { tag: { name: Faker::Commerce.color } } }

      it_behaves_like 'create.js'
    end

    context 'with invalid params' do
      let(:params) { { tag: { name: '' } } }

      it_behaves_like 'new.js'
    end
  end

  describe 'GET /edit.js' do
    let(:tag) { create :tag }

    before { get "/backoffice/tags/#{ tag.id }/edit", xhr: true }

    it_behaves_like 'edit.js'
  end

  describe 'PATCH /update.js' do
    let(:tag) { create :tag }

    before { patch "/backoffice/tags/#{ tag.id }", params:, xhr: true }

    context 'with valid params' do
      let(:params) { { tag: { name: SecureRandom.uuid } } }

      it_behaves_like 'update.js'
    end

    context 'with invalid params' do
      let(:params) { { tag: { name: '' } } }

      it_behaves_like 'edit.js'
    end
  end
end
