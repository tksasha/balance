# frozen_string_literal: true

RSpec.describe 'Categories', type: :request do
  it_behaves_like 'new', uri: '/categories/new.js'

  it_behaves_like 'edit', uri: '/categories/49/edit.js' do
    let(:category) { stub_model Category }

    before { expect(Category).to receive(:find).with('49').and_return(category) }
  end
end
