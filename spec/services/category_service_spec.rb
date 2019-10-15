# frozen_string_literal: true

RSpec.describe CategoryService do
  let :params do
    {
      name: 'Їжа',
      income: true,
      visible: true
    }
  end

  subject { described_class.new params }

  it { should delegate_method(:errors).to(:category) }

  describe '#slug' do
    its(:slug) { should eq 'yizha' }

    context do
      let(:params) { {} }

      its(:slug) { should be_nil }
    end

    context do
      let(:params) { { name: 'Біла Церква' } }

      its(:slug) { should eq 'bila-tserkva' }
    end
  end

  describe '#save' do
    let(:category) { stub_model Category }

    before { expect(Category).to receive(:new).with(params.merge(slug: 'yizha')).and_return(category) }

    context do
      before { expect(category).to receive(:save).and_return(true) }

      its(:save) { should eq true }
    end

    context do
      before { expect(category).to receive(:save).and_return(false) }

      its(:save) { should eq false }
    end
  end
end
