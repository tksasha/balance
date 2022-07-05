# frozen_string_literal: true

RSpec.describe ItemDecorator do
  subject { item.decorate }

  describe '#date' do
    context 'when #date is present' do
      let(:item) { build(:item, date: '2016-10-30') }

      its(:date) { should eq '30.10.2016' }
    end

    context 'when #date is not present' do
      let(:item) { build(:item, date: nil) }

      its(:date) { should be_nil }
    end
  end

  describe '#description' do
    context 'when #description contains one tag' do
      let(:item) { build(:item, description: '[First Tag] some description') }

      its(:description) { should eq '<div class="tag">First Tag</div> some description' }
    end

    context 'when #description contains two tags' do
      let(:item) { build(:item, description: '[First Tag] [Second Tag] some description') }

      its(:description) do
        should eq '<div class="tag">First Tag</div> <div class="tag">Second Tag</div> some description'
      end
    end

    context 'when #description does not contain any tags' do
      let(:item) { build(:item, description: 'some description') }

      its(:description) { should eq 'some description' }
    end

    context 'when #description is nil' do
      let(:item) { build(:item, description: nil) }

      its(:description) { should be_nil }
    end

    context "when #description contains `'`" do
      let(:item) { build(:item, description: "[імперія м'яса] ковбаски") }

      its(:description) { should eq %q(<div class="tag">імперія м'яса</div> ковбаски) }
    end

    context 'when #description contains `-`' do
      let(:item) { build(:item, description: '[а-банк] відсотки за депозитом') }

      its(:description) { should eq '<div class="tag">а-банк</div> відсотки за депозитом' }
    end
  end
end
