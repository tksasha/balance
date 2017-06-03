require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do
  it { should be_an ActsAsRESTController }

  describe '#collection' do
    before do
      #
      # Category.visible.order(:income)
      #
      expect(Category).to receive(:visible) do
        double.tap { |a| expect(a).to receive(:order).with(:income).and_return(:collection) }
      end
    end

    its(:collection) { should eq :collection }
  end

  it_behaves_like :edit

  describe '#resource_params' do
    before { expect(subject).to receive(:params).and_return(acp category: { name: 'Drinks', income: true }) }

    its(:resource_params) { should eq permit! name: 'Drinks', income: true }
  end

  it_behaves_like :update do
    let(:success) { -> { render_template :update } }

    let(:failure) { -> { render_template :errors } }
  end

  it_behaves_like :destroy do
    let(:success) { -> { render_template :destroy } }
  end
end
