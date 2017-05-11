require 'rails_helper'

RSpec.describe ItemsController, type: :controller do
  describe '#items' do
    let(:item) { stub_model Item }

    let(:date) { Date.today }

    let(:date_range) { date.beginning_of_month..date.end_of_month }

    before do
      #
      # Item.search(date_range, nil).includes(:category)
      #
      expect(Item).to receive(:search).with(date_range, nil) do
        double.tap do |a|
          expect(a).to receive(:includes).with(:category)
        end
      end
    end

    it { expect { subject.send :items, date_range }.to_not raise_error }
  end

  describe '#resource' do
    before { expect(subject).to receive(:params).and_return({ id: 43 }) }

    before { expect(Item).to receive(:find).with(43).and_return(:resource) }

    its(:resource) { should eq :resource }
  end

  it_behaves_like :index

  it_behaves_like :index do
    before { @format = :js }
  end

  it_behaves_like :create do
    let(:resource) { stub_model Item }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  it_behaves_like :edit

  it_behaves_like :update do
    let(:resource) { stub_model Item }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :errors } }
  end

  it_behaves_like :destroy do
    let(:success) { -> { should render_template :destroy } }
  end
end
