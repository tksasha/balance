# frozen_string_literal: true

RSpec.describe 'Backoffice::Tags', type: :request do
  describe 'GET /index.js' do
    let(:category) { create :category }

    before { create_list :tag, 2, category: category }

    before { get "/backoffice/categories/#{ category.id }/tags", xhr: true }

    it_behaves_like 'index.js'
  end
end
