# frozen_string_literal: true

RSpec.describe 'Tags', type: :request do
  let(:headers) { { accept: 'application/json' } }

  describe 'GET /index.json' do
    let(:category) { create :category }

    before { create_list :tag, 2, category: category }

    before { get "/categories/#{ category.id }/tags", headers: headers }

    it_behaves_like 'index.json'
  end
end
